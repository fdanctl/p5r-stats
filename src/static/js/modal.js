function closeModal() {
  document.body.classList.remove("modal-open");
  document.getElementById("modal-root").lastElementChild.remove()
}

document.addEventListener("keydown", (e) => {
  if (e.key === "Escape") {
    closeModal();
  }
});

document.addEventListener("htmx:confirm", (e) => {
  console.log(e)
  console.log("question", e.detail.question);
  if (!e.detail.question) return;

  // This will prevent the request from being issued to later manually issue it
  e.preventDefault();

  confirmModal({
    message: e.detail.question,
  }).then(function(result) {
    if (result) {
      e.detail.issueRequest(true); // true to skip the built-in window.confirm()
    }
  });
});

function confirmModal({
  title = "Confirm",
  message = "Are you sure?",
  acceptText = "Yes",
  refuseText = "No",
} = {}) {
  return new Promise((resolve) => {
    const modal = document.createElement("div");

    modal.innerHTML = `
<div
  class="modal-overlay"
  tabindex="-1"
  onclick="event.stopPropagation(); closeModal()"
  onload="document.body.classList.add('modal-open'); this.focus()"
>
  <div
    class="modal fit"
    onclick="event.stopPropagation()"
  >
    <h3>${title}</h3>
    <div>
      <h3>${message}</h3>
      <p class="text-subtitle">This can't be undone</p>
      <div class="flex justify-end items-center">
        <button class="btn-ghost refuse">${refuseText}</button>
        <button class="btn-ghost accept">${acceptText}</button>
      </div>
    </div>
  </div>
</div>
`;
    document.getElementById("modal-root").appendChild(modal);

    modal.querySelector(".accept").onclick = () => {
      modal.remove();
      resolve(true);
    };

    modal.querySelector(".refuse").onclick = () => {
      modal.remove();
      resolve(false);
    };
  });
}
