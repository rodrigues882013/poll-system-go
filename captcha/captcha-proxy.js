#!/usr/bin/nodejs
var http = require('http');

var latestCaptchaJson = JSON.stringify({
 'q':'Is ice hot or cold?',
 'a':['75e52a0ecfafeda17a34fc60111c1f0b'] });
 
var apiUrl = 'http://api.textcaptcha.com/mirror.json';

var refreshCaptchaEveryXSeconds = 5;

var app = http.createServer(function(req,res){
    res.setHeader('Content-Type','application/json');
    res.setHeader('Access-Control-Allow-Origin', '*');
    res.setHeader('Access-Control-Allow-Methods', 'GET, POST');
    res.setHeader('Access-Control-Allow-Headers', 'X-Requested-With,content-type, Authorization');
    res.end(latestCaptchaJson);
});


app.listen(5000);
console.log("Proxy is running")

continuallyRefreshCaptcha();
 
function continuallyRefreshCaptcha(){
 requestNewCaptcha(function(err,captcha){
  setTimeout(continuallyRefreshCaptcha,
             refreshCaptchaEveryXSeconds*1000);
  if(err) return console.log(err); 
  latestCaptchaJson = captcha;
 });
}
 
function requestNewCaptcha(next){
 http.get(apiUrl,function(res){
  var body = '';
  if(res.statusCode!=200){
   return next(new Error('status code '+res.statusCode),null);
  }
  res.on('data',function(chunk){ body+=chunk; });
  res.on('end',function(){
   next(null,body);
  });
 }).on('error',next);
}