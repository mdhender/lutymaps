import * as THREE from 'three';

function main() {
	const systems = [
		{x: 20, y:-4, z:28, kind: "Yellow Main Sequence"},
		{x: 18, y:-3, z:17, kind: "Dense Dust Cloud"},
		{x: 18, y:-2, z:17, kind: "Medium Dust Cloud"},
		{x: 19, y: 7, z:-3, kind: "Blue Super Giant"},
		{x: 10, y:-5, z: 0, kind: "Blue Super Giant"},
		{x:-18, y: 7, z:-2, kind: "Yellow Main Sequence"}
	];

	const canvas = document.querySelector('#c');
	const renderer = new THREE.WebGLRenderer({canvas});

	const fov = 75;
	const aspect = 640 / 480; //2;  // the canvas default
	const near = 1;
	const far = 65;
	const camera = new THREE.PerspectiveCamera(fov, aspect, near, far);
	// camera.position.x = -10;
	// camera.position.y = -10;
	camera.position.z = 49;

	const scene = new THREE.Scene();
	const color = 0xFFFFFF;
	const intensity = 1;
	const light = new THREE.DirectionalLight(color, intensity);
	// light.position.set(-1, 2, 4);
	light.position.set(0, 0, 49);
	scene.add(light);

	const boxWidth = 1;
	const boxHeight = 1;
	const boxDepth = 1;
	const geometry = new THREE.BoxGeometry(boxWidth, boxHeight, boxDepth);

	function makeInstance(geometry, color, id) {
		if (systems[id].kind === "Yellow Main Sequence") {
			color = 0x44aa88;
		} else if (systems[id].kind === "Blue Super Giant") {
			color = 0x8844aa;
		} else if (systems[id].kind === "Light Dust Cloud") {
			color = 0x338866;
		} else if (systems[id].kind === "Medium Dust Cloud") {
			color = 0xaa8844;
		} else if (systems[id].kind === "Dense Dust Cloud") {
			color = 0xbb6688;
		} else {
			color = 0x445566;
		}
		const material = new THREE.MeshPhongMaterial({color});

		const cube = new THREE.Mesh(geometry, material);
		scene.add(cube);

		cube.position.x = systems[id].x;
		cube.position.y = systems[id].y;
		cube.position.z = systems[id].z;

		return cube;
	}

	const cubes = [
		makeInstance(geometry, 0x44aa88, 0),
		makeInstance(geometry, 0x8844aa, 1),
		makeInstance(geometry, 0xaa8844, 2),
		makeInstance(geometry, 0xbb6688, 3),
		makeInstance(geometry, 0x44aa88, 4),
		makeInstance(geometry, 0x44aa88, 5),
	];

	const elem = document.querySelector('#screenshot');
	elem.addEventListener('click', () => {
		drawScene(renderer, scene, camera);
		canvas.toBlob((blob) => {
			saveBlob(blob, `screencapture-${canvas.width}x${canvas.height}.png`);
		});
	});

	const saveBlob = (function() {
		const a = document.createElement('a');
		document.body.appendChild(a);
		a.style.display = 'none';
		return function saveData(blob, fileName) {
			a.href = window.URL.createObjectURL(blob);
			a.download = fileName;
			a.click();
		};
	}());

	function render(time) {
		time *= 0.001;  // convert time to seconds

		cubes.forEach((cube, ndx) => {
			const speed = 1 + ndx * .1;
			const rot = time * speed;
			cube.rotation.x = rot;
			cube.rotation.y = rot;
		});

		renderer.render(scene, camera);

		requestAnimationFrame(render);
	}

	requestAnimationFrame(render);
}

function drawScene(renderer, scene, camera) {
	renderer.render(scene, camera);
}

main();
