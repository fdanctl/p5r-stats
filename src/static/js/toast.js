const randomIntBetween = (min, max) => {
  return Math.floor(Math.random() * (max - min + 1)) + min;
};
const container = document.getElementById("toast-container");

// success, warning, error
function sendToast(type, msg) {
  const toast = document.createElement("div");
  toast.classList.add("toast");
  toast.classList.add(type);
  toast.style.clipPath = `polygon(
  ${randomIntBetween(0, 10)}px 0,
  calc(100% - ${randomIntBetween(0, 10)}px) 0,
  calc(100% - ${randomIntBetween(0, 10)}px) 100%,
  ${randomIntBetween(0, 10)}px 100%
)`;
  toast.addEventListener("animationend", () => {
    toast.remove();
  });

  const p = document.createElement("p");
  p.innerText = msg;

  toast.append(p);
  container.appendChild(toast);

  // const duration = 5000;
  // let elapsed = 0;
  // let intervalId;
  //
  // const finish = () => {
  //   toast.classList.add("transition-out");
  //   clearInterval(intervalId);
  //   toast.classList.remove("visible");
  //   toast.onmouseenter = null;
  //   toast.onmouseleave = null;
  //   setTimeout(() => {
  //     elapsed = duration;
  //     toast.remove();
  //   }, 500);
  // };
  //
  // const start = () => {
  //   setTimeout(() => {
  //     toast.classList.add("visible");
  //   }, 50);
  //   intervalId = setInterval(() => {
  //     elapsed += 10;
  //
  //     if (elapsed >= duration) {
  //       finish();
  //     }
  //   }, 10);
  // };
  //
  // const pause = () => {
  //   clearInterval(intervalId);
  // };
  //
  // start();
  // toast.onmouseenter = pause;
  // toast.onmouseleave = start;
  // toast.onclick = finish;
}

document.body.addEventListener("htmx:responseError", (e) => {
  const msg = e.detail.xhr.response.trim().split("\n");
  for (let i = msg.length - 1; i >= 0; i--) {
    sendToast("error", msg[i]);
  }
  // sendToast("error", e.detail.xhr.response);
});
