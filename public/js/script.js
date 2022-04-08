document.addEventListener("DOMContentLoaded", function() {

var addMinimizeAction       = document.getElementById("addMinimizeAction");
var restoreDefaultAction    = document.getElementById("restoreDefaultAction");
var mainAction              = document.getElementById("mainAction");
var errorWithClosePosition  = document.getElementById("errorWithClosePosition");
var logoImg                 = document.getElementById("logoImg");

    // Asynchronous loading to retrieve different variables as soon as possible
    (async () => {
      try {

        // Check close button position is left
        if (mainAction && errorWithClosePosition) {
            let res = await checkCloseButtonPosition()
            if (!res) {
                mainAction.classList.add("hide");
                errorWithClosePosition.classList.remove("hide");
            }
        }

        // Get actual buttons layout
        if (addMinimizeAction && restoreDefaultAction) {
            let res = await getButtonsLayout();
            if (res) {
                restoreDefaultAction.classList.remove("hide");
                logoImg.classList.remove("opacity-2");
            } else {
                addMinimizeAction.classList.remove("hide");
                logoImg.classList.add("opacity-2");
            }
        }

      }
      catch (e) {
        console.error(e);
      }
    })();

});


// Launch Go function after user action
function add() {
	let res = addMinimizeButton()
    if (res) {
        addMinimizeAction.classList.add("hide");
        restoreDefaultAction.classList.remove("hide");
        logoImg.classList.remove("opacity-2");
    }
}
function restore() {
	let res = restoreButtons()
    if (res) {
        addMinimizeAction.classList.remove("hide");
        restoreDefaultAction.classList.add("hide");
        logoImg.classList.add("opacity-2");
    }
}
