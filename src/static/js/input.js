document.addEventListener('blur', e => {
  if (e.target.matches('.input')) {
    e.target.classList.add('touched');
  }
}, true);
