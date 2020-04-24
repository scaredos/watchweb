// WatchWeb.dart | Dart Implementation of WatchWeb
// Author - ScaredOS | 4/24/2020
// Usage: dart watchweb.dart <url> <filename>
import 'dart:io';
import 'dart:convert';
import 'package:intl/intl.dart';
import 'package:http/http.dart' as http;

main(List<String> args) async {
  int x = 1;
  var url = args[0];
  var fil = args[1];
  var file = new File(fil);
  var date = new DateFormat('yyyy-MM-dd hh:mm:ss');
  var sink = file.openWrite();
  while(x < 2) {
    var res = await http.get(url);
    //print('[${res.statusCode}] Response code on ($url) - ${date.format(new DateTime.now())}'); // Uncomment to print to console
    sink.write('[${res.statusCode}] Response code on ($url) - ${date.format(new DateTime.now())}\n');
    sleep(const Duration(seconds:1));
  }
}
