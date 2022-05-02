const express = require('express');
const spawn = require('child_process').execFile;
const path = require('path');
const app = express();

app.use(express.static('./'));

app.get('/', (req, res) => {
	res.render('index.html');
});

app.get('/getmeme', (req, res) => { 
	let name = Date.now().toString();

	randomTemplate(name, () => {
		res.sendFile(path.join(__dirname, `memes/generated/${name}.png`));
	});
});

app.listen(3000, () => {
	console.log('Server running on port 3000');
});

async function randomTemplate(name, cb) {
	const goproc = spawn('meme_generator.exe', [`memes/templates/`, 'memes/generated/', name], (err, stdout, stderr) => {
		if (err) throw err;
		if (stderr) console.log("STDERR: " + stderr);
		if (stdout) console.log("STDOUT: " + stdout);
	});

	goproc.on('exit', (code) => { 
		if (code != 0) console.log("Exit code: " + code);
		cb();
	});
}