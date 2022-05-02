onload = () => {
	let imgEl = document.getElementById("meme");

	fetch("http://localhost:3000/getmeme")
		.then((response) => response.blob())
		.then((imgBlob) => {
			const imgObjURL = URL.createObjectURL(imgBlob);
			imgEl.src = imgObjURL;
		});
};
