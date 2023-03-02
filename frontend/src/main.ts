import { Terminal } from "xterm";

const term = new Terminal();

term.open(document.getElementById("terminal")!);

//@ts-ignore
window.__WRITE_TERMINAL = (data: string) => {
	term.write(data);
};
