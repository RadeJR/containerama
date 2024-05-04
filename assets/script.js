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
