import { AmnesiaClient } from '@amnesia-js/core';
import cron from 'node-cron';
import tcpPortUsed from 'tcp-port-used';
import express from 'express';
import cors from 'cors';

const DETERMINISTIC_PORT_RANGE = 2000;
const PORT = 4422;
const LBED_READ_PORT = 4224;
const read_workers = [];

const app = express();
app.use(cors());
app.use(express.json());

let READ_HOST = '127.0.0.1';
let WRITE_HOST = READ_HOST;
let LOADBALANCED_READ_HOST = READ_HOST;
// if (process.env.NODE_ENV === 'production') {
//   READ_HOST = 'read_replica';
//   WRITE_HOST = 'write_replica';
//   LOADBALANCED_READ_HOST = 'loadbalancer';
// }

const write_worker = new AmnesiaClient();
await write_worker.connect({ port: 4223 });

for (let i = 0; i <= DETERMINISTIC_PORT_RANGE; i++) {
  const port = 3000 - i;
  try {
    // console.log('Checking port', port);
    await tcpPortUsed.check(port)
      .then(async (inUse) => {
        // console.log(port)
        if (inUse) {
          try {
            const client = new AmnesiaClient();
            client.connect({ port }).then(() => {
              read_workers.push(client);
              console.log(`Found port: ${port}`)
            }).catch((err) => {
              console.log(err);
            });
          } catch (e) {
            console.log(e);
          }
        }
      }, () => {
      });
  } catch (e) {

  }
}

console.log('🔗 All workers connected!')

let fanout_queue = [];
let query_replay = [];

cron.schedule('*/10 * * * * *', () => {
  console.log("SCHEDULE: FANOUT")
  query_replay.forEach((query) => {
    read_workers.forEach((worker) => {
      worker.query(query).then();
    });
  })
  fanout_queue.forEach(({ key, value }) => {
    read_workers.forEach((worker) => {
      worker.set(key, value).then();
    })
  });
  query_replay = [];
  fanout_queue = [];
});

app.post('/cmd', async (req, res) => {
  const { query } = req.body;
  console.log('📡 CMD', req.body);
  query_replay.push(query);
  if (query.startsWith('SET')) {
    query_replay.push(query);
    const response = await write_worker.query(query);
    console.log(response)
    res.send({ ok: true, response });
  } else {
    let lbedClient = new AmnesiaClient();
    await lbedClient.connect({ port: LBED_READ_PORT, host: LOADBALANCED_READ_HOST });
    const response = await lbedClient.query(query);
    lbedClient.destroy();
    console.log(response)
    res.send({ ok: true, data: response });
  }

})

app.post('/query', async (req, res) => {
  const { key, value } = req.body;
  fanout_queue.push({
    key,
    value,
  });

  const response = await write_worker.set(key, value);

  res.send(response);
});

app.get('/query', async (req, res) => {
  const { key } = req.query;

  console.log(key)
  let lbedClient = new AmnesiaClient();
  await lbedClient.connect({ port: LBED_READ_PORT, host: LOADBALANCED_READ_HOST });
  const response = await lbedClient.get(key);
  lbedClient.destroy();
  res.send(response);
});

app.listen(PORT)