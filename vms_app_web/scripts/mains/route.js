const users = [
    { type: "task", value: "Routes" },
    { type: "button", value: "+ Add Route" },
];

const cols = [
    { type: "ID", value: "ID" },
    { type: "medium", value: "Start Location" },
    { type: "medium", value: "End Location" },
    { type: "short", value: "Time" },
    { type: "long", value: "Distance" },
    { type: "status", value: "Status" },
    { type: "action", value: "Action" },
];

const $top = document.querySelector(".top");
const $columns = document.querySelector(".columns");

users.forEach((block) => {
    let toph = "";
    if (block.type === "task") {
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
    } else if (block.type === "action") {
        html = action(block);
    } else if (block.type === "status") {
        html = stat(block);
    } else if (block.type === "ID") {
        html = `<div style="width: 135px;"class="text">
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

function short(block) {
    return `
    <div style="flex: 1;" class="text">
        <p>${block.value}</p>
    </div>
`;
}

function stat(block) {
    return `
    <div style="flex: 1.065;" class="text">
        <p>${block.value}</p>
    </div>
`;
}

function medium(block) {
    return `
    <div style="flex: 1.3;" class="text">
        <p>${block.value}</p>
    </div>
`;
}

function long(block) {
    return `
    <div style="flex: 1.2;" class="text">
        <p>${block.value}</p>
    </div>
`;
}

function action(block) {
    return `
    <div style="flex: 0.5;" class="action">
        <p>${block.value}</p>
    </div>
`;
}
