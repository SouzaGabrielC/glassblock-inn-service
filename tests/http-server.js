const http = require("http");

const server = http.createServer(function (req, res) {
  console.log(req.method, req.url);
  req.on("data", (data) => {
    console.log("Data", data.toString());
  });
  res.writeHead(200, { "Content-Type": "application/json" });
  res.end(
    JSON.stringify({
      data: "Hello World!",
    })
  );
});

server.listen(8080);
