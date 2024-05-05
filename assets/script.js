function closePopup() {
  // Get a reference to the original div
  var originalDiv = document.getElementById("popup");

  // Create a new div
  var newDiv = document.createElement("div");
  newDiv.id = "popup";
  newDiv.innerHTML = "";

  // Replace the original div with the new one
  originalDiv.parentNode.replaceChild(newDiv, originalDiv);
}

document.addEventListener("DOMContentLoaded", (event) => {
	document.body.addEventListener('htmx:beforeSwap', function(evt) {
		if (evt.detail.xhr.status === 422 || evt.detail.xhr.status === 400) {
			evt.detail.shouldSwap = true;
			evt.detail.isError = false;
		}
	});
})
