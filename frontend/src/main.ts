import { Terminal } from "xterm";

const term = new Terminal();

let command = "";

term.onKey(({ key, domEvent }) => {
	const isPrintable = !(
		domEvent.altKey ||
		domEvent.ctrlKey ||
		domEvent.metaKey
	);

	if (domEvent.key === "Enter") {
		//@ts-ignore
		window.__WRITE_PTY(command);
		term.writeln("");
		command = "";
		return;
	}

	if (isPrintable) {
		term.write(key);
		command += key;
	}
});

term.open(document.getElementById("terminal")!);

//@ts-ignore
window.__WRITE_TERMINAL = (data: { result: string }) => {
	term.write(data.result);
};
