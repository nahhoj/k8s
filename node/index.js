const http=require("http");
const express=require("express");
const app=express();
const os=require('os');

const port=8080;

let server=http.createServer(app)
    .listen(port,()=>{
        console.log(`Server is running ${port}`);
    });

app.get("/",(req,res)=>{
    let ip;
    for(let i in os.networkInterfaces()){
        for (let x in os.networkInterfaces()[i]){
             if (os.networkInterfaces()[i][x].family === 'IPv4' && os.networkInterfaces()[i][x].mac!=='00:00:00:00:00:00'){
                 ip=os.networkInterfaces()[i][x].address;
             }
        }
     }
    res.status(200);
    res.setHeader('Content-type','text/html; charset=UTF-8');
    res.send(`<h2>Hello World node with kubernetes and microservice</h2>
            <h3>Platform: ${os.platform()}</h3>
            <h3>Hostmame: ${os.hostname()}</h3>
            <h3>IP: ${ip}</h3>`);
});