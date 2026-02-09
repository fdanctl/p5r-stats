const allValues = ["knowledge", "guts", "proficiency", "kindness", "charm"];
function customOnChange(minus) {
  const selects = document.getElementsByClassName("select-stat");
  let usedValues = [];
  for (let i = 0; i < selects.length; i++) {
    const value = selects[i].querySelector("select").value;
    console.log(value);
    usedValues = usedValues.concat(value);
  }

  console.log(usedValues);
  if (minus) {
    const idx = usedValues.findIndex((e) => e === minus);
    if (idx > 0) {
      usedValues = usedValues
        .slice(0, idx)
        .concat(...usedValues.slice(idx + 1));
    }
  }
  console.log(usedValues);

  let possibleValues = allValues.filter((e) => !usedValues.includes(e));
  console.log(possibleValues);

  for (let i = 0; i < selects.length; i++) {
    const select = selects[i].querySelector("select");
    const value = select.value;
    select.innerHTML = "";

    const starter = document.createElement("option");
    starter.value = "";
    starter.disabled = true;
    starter.selected = true;
    starter.hidden = true;
    starter.innerText = "Add stat";
    select.appendChild(starter);

    const values = [value].concat(...[possibleValues]);
    console.log(values);
    for (let j = 0; j < values.length; j++) {
      const v = values[j];
      const opt = document.createElement("option");
      opt.value = values[j];
      opt.innerText = v.charAt(0).toUpperCase() + v.slice(1).toLowerCase();

      if (j === 0) {
        opt.selected = true;
      }

      if (v === "") {
        opt.innerText = "Add stat";
        opt.disabled = true;
        opt.hidden = true;
      }
      select.appendChild(opt);
    }
  }
}

function customDelete(element) {
  console.log(element);
  const selects = [...document.getElementsByClassName("select-stat")].map((e) =>
    e.querySelector("select"),
  );
  const thisSelect = element.parentElement.querySelector("select");
  console.log(selects);

  if (thisSelect.value !== "") {
    console.log(selects.map((e) => e.value));
    if (selects.every((e) => e.value !== "")) {
      const clone = element.parentElement.cloneNode(true)
      clone.querySelector("select").value = ""
      console.log(clone)
      document.getElementById("stats-selects").appendChild(clone)
    }
    element.parentElement.remove();
  }
}
