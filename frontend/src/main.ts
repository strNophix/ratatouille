import { Terminal } from "xterm";

const term = new Terminal({
	fontFamily:
		'"Fira Code", courier-new, courier, monospace, "Powerline Extra Symbols"',
});

let line = "";

term.onKey(({ key, domEvent }) => {
	const printable = !(domEvent.altKey || domEvent.ctrlKey || domEvent.metaKey);
	if (domEvent.key === "Enter") {
		//@ts-ignore
		window.__WRITE_PTY(line);
		term.writeln("");
		line = "";
		return;
	}

	if (printable) {
		term.write(key);
		line += key;
	}
});

term.open(document.getElementById("terminal")!);

//@ts-ignore
window.__WRITE_TERMINAL = (data: { result: string }) => {
	term.write(data.result);
};
