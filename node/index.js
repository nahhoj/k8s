'use strict'

const http=require("http");
const express=require("express");
const app=express();
const os=require('os');
const path=require('path')

const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');


const PROTO_PATH=path.resolve(__dirname,"../proto/message.proto");

const portWebServer=8080;
const portgRPCServer=8081;

let packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });

let servMessage=grpc.loadPackageDefinition(packageDefinition).main;
const gRPCserver=new grpc.Server();

gRPCserver.addService(servMessage.Message.service,{
    SendMessage:function(call,callback){
        callback(null,{iam:os.hostname()});
    }
});

let client = new servMessage.Message('go.johan.lcl:3001',grpc.credentials.createInsecure());

gRPCserver.bind(`0.0.0.0:${portgRPCServer}`,grpc.ServerCredentials.createInsecure());
gRPCserver.start();

let webServer=http.createServer(app)
    .listen(portWebServer,()=>{
        console.log(`Server is running ${portWebServer}`);
    });

app.get("/",(req,res)=>{
    let resMessage;
    client.SendMessage({who:'dsdsd'},function(err,resp){
       if (err) resMessage=err.details
       else resMessage=resp.iam
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
       res.send(`<h2>Hello World <span style=\"color:red;\">NODE</span> with kubernetes and microservice</h2>
               <h3>Platform: ${os.platform()}</h3>
               <h3>Hostmame: ${os.hostname()}</h3>
               <h3>IP: ${ip}</h3>
               <br/>
               <h3>Message from: ${resMessage}</h3>`);
    });
});