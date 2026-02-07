document.addEventListener(
  "blur",
  (e) => {
    if (e.target.matches(".input")) {
      e.target.classList.toggle("error", !e.target.validity.valid);
      const parent = e.target.parentElement;
      console.log(parent.classList);
      if (parent.classList.contains("input-group")) {
        const errorSpan = parent.querySelector("span");

        const isRequired = e.target.required;
        if (isRequired) {
          const field = e.target.name;
          errorSpan.innerText = 
            `${field[0].toUpperCase() + field.substring(1)} is required`;
          return;
        }
      }
    }
  },
  true,
);
