const users = [
    { type: "vehicle", value: "Auctions" },
    { type: "button", value: "+ Add Auction" },
];

const cols = [
    { type: "ID", value: "ID" },
    { type: "long", value: "VIN" },
    { type: "lpn", value: "Auction name" },
    { type: "short", value: "Start Price" },
    { type: "short", value: "Current Price" },
    { type: "status", value: "Status" },
    { type: "action", value: "Action" },
];

const $top = document.querySelector(".top");
const $columns = document.querySelector(".columns");

users.forEach((block) => {
    let toph = "";
    if (block.type === "vehicle") {
        toph = title(block);
    } else if (block.type === "button") {
        toph = button(block);
    }
    $top.insertAdjacentHTML("beforeend", toph);
});

cols.forEach((block) => {
    let html = "";

    if (block.type === "short") {
        html = short(block);
    } else if (block.type === "medium") {
        console.log(block.value);
        html = medium(block);
    } else if (block.type === "long") {
        html = long(block);
    } else if (block.type === "make") {
        html = make(block);
    } else if (block.type === "action") {
        html = action(block);
    } else if (block.type === "miles") {
        html = miles(block);
    } else if (block.type === "status") {
        html = stat(block);
    } else if (block.type === "lpn") {
        html = lpn(block);
    } else if (block.type === "ID") {
        html = `<div style="width: 130px;"class="text">
                    <p>${block.value}</p>
            </div>`;
    }
    $columns.insertAdjacentHTML("beforeend", html);
});

function title(block) {
    return `
    <div style="flex: 2;"class="title">
        <h1>${block.value}</h1>
    </div>
`;
}
function button(block) {
    return `
    <button type="button" class="add-button">
        <p>${block.value}</p>
    </div>
`;
}

function make(block) {
    return `
    <div style="flex: 0.9;" class="text">
        <p>${block.value}</p>
    </div>
`;
}

function stat(block) {
    return `
    <div style="flex: 0.65;" class="text">
        <p>${block.value}</p>
    </div>
`;
}
function lpn(block) {
    return `
    <div style="flex: 1; margin-left:20px;" class="text">
        <p>${block.value}</p>
    </div>
`;
}

function short(block) {
    return `
    <div style="flex: 0.8;" class="text">
        <p>${block.value}</p>
    </div>
`;
}
function miles(block) {
    return `
    <div style="flex: 0.6; font-size:17px;" class="text">
        <p>${block.value}</p>
    </div>
`;
}

function medium(block) {
    return `
    <div style="flex: 1;" class="text">
        <p>${block.value}</p>
    </div>
`;
}

function long(block) {
    return `
    <div style="flex: 0.9;" class="text">
        <p>${block.value}</p>
    </div>
`;
}

function action(block) {
    return `
    <div style="flex: 0.57;" class="action">
        <p>${block.value}</p>
    </div>
`;
}
