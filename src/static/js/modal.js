function closeModal() {
  document.body.classList.remove("modal-open")
  document.getElementById("modal-root").innerHTML = "";
}

document.addEventListener("keydown", (e) => {
  if (e.key === "Escape") {
    closeModal();
  }
});
