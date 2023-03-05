import { Terminal } from "xterm";

const term = new Terminal();

term.onKey(({ key }) => {
	//@ts-ignore
	window.__WRITE_PTY(key);
});

term.open(document.getElementById("terminal")!);

//@ts-ignore
window.__WRITE_TERMINAL = (data: { result: string }) => {
	term.write(data.result);
};
