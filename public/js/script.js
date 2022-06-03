document.addEventListener("DOMContentLoaded", function() {

var addMinimizeAction       = document.getElementById("addMinimizeAction");
var restoreDefaultAction    = document.getElementById("restoreDefaultAction");
var elementaryIcons         = document.getElementById("elementaryIcons");
var eosMinimizeIcon         = document.getElementById("eosMinimizeIcon");

var macIcons                = document.getElementById("macIcons");
var selectMacButtons        = document.getElementById("selectMacButtons");
var restoreDefaultMac       = document.getElementById("restoreDefaultMac");

var winIcons                = document.getElementById("winIcons");
var selectWinButtons        = document.getElementById("selectWinButtons");
var restoreDefaultWin       = document.getElementById("restoreDefaultWin");

var wtfIcons                = document.getElementById("wtfIcons");
var selectWtfButtons        = document.getElementById("selectWtfButtons");
var restoreDefaultWtf       = document.getElementById("restoreDefaultWtf");

    // Asynchronous loading to retrieve different variables as soon as possible
    (async () => {
      try {

        // Check close button position is left
        if (elementaryIcons && macIcons) {

            let style = await checkButtonsStyle();
            if (style == 'eos') {
                addMinimizeAction.classList.remove("hide");
                eosMinimizeIcon.classList.add("opacity-2");
                elementaryIcons.classList.add("buttons-bar-selected");
                selectMacButtons.classList.remove("hide");
                selectWinButtons.classList.remove("hide");
                selectWtfButtons.classList.remove("hide");

            } else if (style == 'eos+min') {
                restoreDefaultAction.classList.remove("hide");
                eosMinimizeIcon.classList.remove("opacity-2");
                elementaryIcons.classList.add("buttons-bar-selected");
                selectMacButtons.classList.remove("hide");
                selectWinButtons.classList.remove("hide");
                selectWtfButtons.classList.remove("hide");

            } else if (style == 'mac') {
                macIcons.classList.add("buttons-bar-selected");
                restoreDefaultMac.classList.remove("hide");

            } else if (style == 'win') {
                winIcons.classList.add("buttons-bar-selected");
                restoreDefaultWin.classList.remove("hide");

            } else if (style == 'wtf') {
                wtfIcons.classList.add("buttons-bar-selected");
                restoreDefaultWtf.classList.remove("hide");

            }
        }

      }
      catch (e) {
        console.error(e);
      }
    })();

});


// Launch Go function after user action
function cleanElements() {
    addMinimizeAction.classList.add("hide");
    restoreDefaultAction.classList.add("hide");
    eosMinimizeIcon.classList.add("opacity-2");
    elementaryIcons.classList.remove("buttons-bar-selected");

    selectMacButtons.classList.add("hide");
    restoreDefaultMac.classList.add("hide");
    macIcons.classList.remove("buttons-bar-selected");

    selectWinButtons.classList.add("hide");
    restoreDefaultWin.classList.add("hide");
    winIcons.classList.remove("buttons-bar-selected");

    selectWtfButtons.classList.add("hide");
    restoreDefaultWtf.classList.add("hide");
    wtfIcons.classList.remove("buttons-bar-selected");
}

function setDefault() {
    elementaryIcons.classList.add("buttons-bar-selected");

    selectMacButtons.classList.remove("hide");
    selectWinButtons.classList.remove("hide");
    selectWtfButtons.classList.remove("hide");
}

function addEosMinimize() {
	let res = addMinimizeButton()
    if (res) {
        cleanElements();
        setDefault();

        restoreDefaultAction.classList.remove("hide");
        eosMinimizeIcon.classList.remove("opacity-2");
    }
}

function setEosDefaultStyle() {
	let res = restoreButtons()
    if (res) {
        cleanElements();
        setDefault();

        addMinimizeAction.classList.remove("hide");
    }
}

function setMacStyle() {
	let res = applyMacButtons()
    if (res) {
        cleanElements();

        restoreDefaultMac.classList.remove("hide");
        macIcons.classList.add("buttons-bar-selected");
    }
}

function setWinStyle() {
	let res = applyWinButtons()
    if (res) {
        cleanElements();

        restoreDefaultWin.classList.remove("hide");
        winIcons.classList.add("buttons-bar-selected");
    }
}

function setWtfStyle() {
	let res = applyWtfButtons()
    if (res) {
        cleanElements();

        restoreDefaultWtf.classList.remove("hide");
        wtfIcons.classList.add("buttons-bar-selected");
    }
}
