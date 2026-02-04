function openFirstChildFileInput(element) {
  const fileInput = element.getElementsByTagName("input")[0];
  fileInput.click();
}

let pfpOGSrc
function start() {
  const pfpImg = document.querySelector("#pfp img");
  pfpOGSrc = pfpImg?.src ?? null;
  console.log(pfpOGSrc);

  const fileInput = document.getElementById("pfpChange");
  fileInput &&
    fileInput.addEventListener("change", () => {
      const file = fileInput.files[0];
      if (!file) return;

      if (file.size > 10 * 1024 * 1024) {
        console.warn("file size to large");
        sendToast("error", "file size to large");
        return;
      }

      const url = URL.createObjectURL(file);
      pfpImg.src = url;

      pfpImg.onload = () => URL.revokeObjectURL(url);
    });

  document.getElementById("user-info")?.addEventListener("reset", () => {
    pfpImg.src = pfpOGSrc;
  });
}

start();

document.body.addEventListener("htmx:after-swap", () => {
  !pfpOGSrc && start();
});


