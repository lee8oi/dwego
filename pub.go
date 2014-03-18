// pub.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func init() {
	err := os.Mkdir("./pub/", 0744)
	if err == nil { //dir was created, likely needs files.
		err = ioutil.WriteFile("./pub/style.css", STYLES, 0644)
		if err != nil {
			fmt.Println("Error writing style.css: ", err)
		}
		err = ioutil.WriteFile("./pub/home.html", HTML, 0644)
		if err != nil {
			fmt.Println("Error writing home.html: ", err)
		}
		err = ioutil.WriteFile("./pub/script.js", SCRIPT, 0644)
		if err != nil {
			fmt.Println("Error writing script.js: ", err)
		}
		fmt.Println("pub dir & web files created")
	} else {
		fmt.Println("loading existing pub dir")
	}
	homeTempl = template.Must(template.ParseFiles("pub/home.html"))
	scriptTmpl = template.Must(template.ParseFiles("pub/script.js"))
}

var STYLES = []byte(`
html {
    overflow: hidden;
}

body {
    overflow: hidden;
    padding: 0;
    margin: 0;
    width: 100%;
    height: 100%;
    background: gray;
}

#log {
    background: white;
    margin: 0;
    padding: 0.5em 0.5em 0.5em 0.5em;
    position: absolute;
    top: 0.5em;
    left: 0.5em;
    right: 0.5em;
    bottom: 3em;
    overflow: auto;
}

#form {
    padding: 0 0.5em 0 0.5em;
    margin: 0;
    position: absolute;
    bottom: 1em;
    left: 0px;
    width: 100%;
    overflow: hidden;
}
`)

var HTML = []byte(`
<html>
<head>
<title>Dwego MUD</title>
<script type="text/javascript" src="http://ajax.googleapis.com/ajax/libs/jquery/1.4.2/jquery.min.js"></script>
<script src="http://crypto-js.googlecode.com/svn/tags/3.1.2/build/rollups/sha256.js"></script>
<!--<script>
    var hash = CryptoJS.SHA256("Message");
</script>-->
<script type="text/javascript" src="/scripts"></script>
<link rel="stylesheet" type="text/css" href="/public/style.css">
</head>
<body>
<div id="log"></div>
<form id="form">
    <input type="submit" value="Send" />
    <input type="text" id="msg" size="64"/>
</form>
</body>
</html>
`)

var SCRIPT = []byte(`
$(function() {
	var conn;
	var msg = $("#msg");
	var log = $("#log");
	
	function appendLog(msg) {
		var d = log[0]
		var doScroll = d.scrollTop == d.scrollHeight - d.clientHeight;
		msg.appendTo(log)
		if (doScroll) {
			d.scrollTop = d.scrollHeight - d.clientHeight;
		}
	}
	
	$("#form").submit(function() {
		if (!conn) {
			return false;
		}
		if (!msg.val()) {
			return false;
		}
		conn.send(msg.val());
		msg.val("");
		return false
	});
	
	if (window["WebSocket"]) {
		conn = new WebSocket("ws://{{$}}/ws");
		conn.onclose = function(evt) {
			appendLog($("<div><b>Connection closed.</b></div>"))
		}
		conn.onmessage = function(evt) {
			appendLog($("<div/>").text(evt.data))
		}
	} else {
		appendLog($("<div><b>Your browser does not support WebSockets.</b></div>"))
	}
});
`)
