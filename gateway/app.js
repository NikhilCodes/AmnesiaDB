import { AmnesiaClient } from '@amnesia-js/core';
import cron from 'node-cron';
import tcpPortUsed from 'tcp-port-used';
import express from 'express';

const DETERMINISTIC_PORT_RANGE = 10000;
const PORT = 4422;
const LBED_READ_PORT = 4224;
const read_workers = [];

const app = express();
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

for (let i = 0; i < DETERMINISTIC_PORT_RANGE; i++) {
  const port = 68200 - i;
  await tcpPortUsed.check(port)
    .then(async (inUse) => {
      // console.log(port)
      if (inUse) {
        console.log(`Found port: ${port}`)
        const client = new AmnesiaClient();
        await client.connect({ port });
        read_workers.push(client);
      }
    }, () => {});
}

console.log('🔗 All workers connected!')

let fanout_queue = [];

cron.schedule('*/10 * * * * *', () => {
  console.log("SCHEDULE: FANOUT")
  fanout_queue.forEach(({ key, value }) => {
    read_workers.forEach((worker) => {
      worker.set(key, value).then();
    })
  });
  fanout_queue = [];
});


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