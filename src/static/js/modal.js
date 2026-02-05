// function toggleModal(id) {
//   const modal = document.getElementById(id)
//   modal.classList.toggle('visible');
//
//   const focusable = modal.querySelectorAll(
//     'button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])'
//   );
//   const first = focusable[0];
//   const last = focusable[focusable.length - 1];
//
//   modal.addEventListener('keydown', (e) => {
//     if (e.key !== 'Tab') return;
//
//     if (e.shiftKey) {
//       // Shift+Tab on first element → go to last
//       if (document.activeElement === first) {
//         e.preventDefault();
//         last.focus();
//       }
//     } else {
//       // Tab on last element → go to first
//       if (document.activeElement === last) {
//         e.preventDefault();
//         first.focus();
//       }
//     }
//   });
//   // start focus on the modal (must have tabindex="-1")
//   modal.querySelector('div:last-child')
// }

function closeModal() {
  document.body.classList.remove("modal-open")
  document.getElementById("modal-root").innerHTML = "";
}

document.addEventListener("keydown", (e) => {
  if (e.key === "Escape") {
    closeModal();
  }
});
