const $sbButtons = document.querySelectorAll(".sb-button");
const currentUrl = window.location.href;

let whichButton = "driver";

if (currentUrl.includes("route.html")) {
    whichButton = "route";
} else if (currentUrl.includes("driver.html")) {
    whichButton = "driver";
} else if (currentUrl.includes("vehicle.html")) {
    whichButton = "vehicle";
} else if (currentUrl.includes("auction.html")) {
    whichButton = "auction";
} else if (currentUrl.includes("fueler.html")) {
    whichButton = "fueler";
} else if (currentUrl.includes("maintainer.html")) {
    whichButton = "maintainer";
} else if (currentUrl.includes("fueling.html")) {
    whichButton = "fueling";
} else if (currentUrl.includes("maintenance.html")) {
    whichButton = "maintenance";
} else if (currentUrl.includes("report.html")) {
    whichButton = "report";
}

$sbButtons.forEach((button) => {
    if (button.id === whichButton) {
        button.style.backgroundColor = "aliceblue";
    }
});
