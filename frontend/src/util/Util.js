export function setTitle(page) {
  document.title = `Todo App | ${page}`;
}

export class Timer {
  constructor() {
    this.timer = null;
  }

  keydown() {
    clearTimeout(this.timer);
  }

  keyup(f, t) {
    clearTimeout(this.timer);
    this.timer = setTimeout(f, t);
  }
}
