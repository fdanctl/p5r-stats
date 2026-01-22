function showNearestTooltip(element) {
  const tooltip = element.getElementsByClassName("tooltip")[0];
  tooltip.classList.add("visible");
}

function hideNearestTooltip(element) {
  const tooltip = element.getElementsByClassName("tooltip")[0];
  tooltip.classList.remove("visible");
}
