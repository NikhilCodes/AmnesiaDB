const net = require('net')
const readline = require("readline").createInterface({
  input: process.stdin,
  output: process.stdout,
});
const options = {
  port: 4224,
};

let client = net.connect(options, () => {
  console.log("connected!");
  client.write("SET a AS v\n");
});

client.on('data' ,(data) => {
    console.log('received data');
    console.log(data.toString());
})
