const pencilBtn = document.getElementById("pfpEditBtn");
const nameH1 = document.querySelector(".editUsername h1");
const nameInput = document.getElementById("nameInput");
const pencilSVG = document.querySelector(".username svg");
const formBtns = document.getElementById("formBtns");
const fileInput = document.getElementById("pfpChange");
const pfpImg = document.querySelector("#pfp img");
let pfpOGSrc = pfpImg.src;

function toggleEditState() {
  pencilBtn.classList.toggle("visible");
  nameH1.classList.toggle("hidden");
  nameInput.classList.toggle("hidden");
  pencilSVG.classList.toggle("visible");
  formBtns.classList.toggle("visible");
}

const pfp = document.getElementById("pfp");

pfp.addEventListener("click", () => {
  !pencilBtn.classList.contains("visible") && toggleEditState();

  fileInput.click();
});

pencilSVG.addEventListener("click", () => {
  if (!pencilSVG.classList.contains("visible")) {
    toggleEditState();
    nameInput.focus();
    const length = nameInput.value.length;
    nameInput.setSelectionRange(length, length);
  }
});

nameH1.addEventListener("click", () => {
  toggleEditState();
  nameInput.focus();
  const length = nameInput.value.length;
  nameInput.setSelectionRange(length, length);
});

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

document
  .querySelector("#formBtns button[type=reset]")
  .addEventListener("click", () => {
    pfpImg.src = pfpOGSrc;
    toggleEditState();
  });

// ======== //

document.getElementById("userInfo").addEventListener("submit", (event) => {
  event.preventDefault();

  const formData = new FormData(event.target);

  for (const [k, v] of formData.entries()) {
    console.log(k, v);
  }

  fetch("/api/user-data", {
    method: "PATCH",
    body: formData,
  });

  const file = fileInput.files[0];
  if (file) {
    pfpOGSrc = URL.createObjectURL(file);
  }

  toggleEditState();
});
