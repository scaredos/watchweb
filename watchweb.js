// WatchWeb.js | NodeJS implementation of WatchWeb
// Author - ScaredOS | 4/16/2020
// Usage: node watchweb.js <url> <filename>
'use strict';
var fs = require('fs');
var https = require('https');
var util = require('util');
var url = process.argv[2];
var filename = process.argv[3];
var data = '';


function get() {
  https.get(url, res => {
    data = util.format("[%s] Response code on (%s) - %s\n", res.statusCode, url, new Date().toISOString().replace(/T/, ' ').replace(/\..+/, ''));
    // console.log(data); // Uncomment this line to log the data to console
    fs.appendFile(filename, data, function (err) {
      if (err) throw err;
    });
  });
}

get();
setInterval(get, 300000);
