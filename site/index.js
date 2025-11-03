"use strict";


const go = new Go();
let wasm, codec;

(async function init() {
	wasm = await WebAssembly.instantiateStreaming(fetch("codec.wasm"), go.importObject)
	go.run(wasm.instance);
	codec = wasm.instance.exports;

	// We are ready; enable the buttons.
	for (const btn of document.getElementsByClassName("submit-btn"))
		btn.removeAttribute("disabled");
})();


function openEncodeTab() {
	document.getElementById("header-btn-encode").className = "active";
	document.getElementById("header-btn-decode").className = "inactive";
	document.getElementById("decode-tab").style.display = "none";
	document.getElementById("encode-tab").style.display = "block";
}

function openDecodeTab() {
	document.getElementById("header-btn-decode").className = "active";
	document.getElementById("header-btn-encode").className = "inactive";
	document.getElementById("encode-tab").style.display = "none";
	document.getElementById("decode-tab").style.display = "block";
}


function codecCall(buf, fn) {
	let inAddr = codec.requestBuffer(buf.length);
	(new Uint8Array(codec.memory.buffer, inAddr, buf.length)).set(buf);

	let outAddr = fn();
	let outLen = codec.getOutLen();

	let err = getCodecErrorMsg();
	if (err !== null) {
		alert(err);
		return;
	}

	return new Uint8Array(codec.memory.buffer, outAddr, outLen);
}


async function generateImage() {
	let file = document.getElementById("encode-in").files[0];

	if (file.size > 2**32 - 1) {
		alert(`File too large! (${file.size} bytes)`);
		return;
	}

	let payload = new Uint8Array(await file.arrayBuffer());
	let pngData = codecCall(payload, codec.encode);

	document.getElementById("encode-out-header").style.display = "inherit";

	let out = document.getElementById("encode-out");
	out.src = URL.createObjectURL(new Blob([pngData]));
}


async function generateFile() {
	let image = document.getElementById("decode-in").files[0];
	if (image === undefined)
		return;

	let imgData = new Uint8Array(await image.arrayBuffer());
	let payload = codecCall(imgData, codec.decode);

	let blob = new Blob([payload], {type: "application/octet-stream"});
	let out = document.getElementById("decode-out");
	out.download = "thing";
	out.href = URL.createObjectURL(blob);
	out.style.display = "inherit";
}
