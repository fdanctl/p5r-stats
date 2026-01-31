function showFirstChildTooltip(element) {
  const tooltip = element.getElementsByClassName("tooltip")[0];
  tooltip.classList.add("visible");
}

function hideFirstChildTooltip(element) {
  const tooltip = element.getElementsByClassName("tooltip")[0];
  tooltip.classList.remove("visible");
}
