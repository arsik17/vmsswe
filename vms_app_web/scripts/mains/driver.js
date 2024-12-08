const users = [
    { type: "driver", value: "Drivers" },
    { type: "button", value: "+ Add Driver" },
];

const cols = [
    { type: "ID", value: "ID" },
    { type: "long", value: "Email" },
    { type: "names", value: "First" },
    { type: "surname", value: "Middle" },
    { type: "surname", value: "Last" },
    { type: "phone", value: "Phone" },
    { type: "medium", value: "Gov. id" },
    { type: "action", value: "Action" },
];

const $top = document.querySelector(".top");
const $columns = document.querySelector(".columns");

users.forEach((block) => {
    let toph = "";
    if (block.type === "driver") {
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
    } else if (block.type === "names") {
        html = names(block);
    } else if (block.type === "phone") {
        html = phone(block);
    } else if (block.type === "surname") {
        html = surname(block);
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

function short(block) {
    return `
    <div style="flex: 0.5;" class="text">
        <p>${block.value}</p>
    </div>
`;
}
function phone(block) {
    return `
    <div style="flex: 1.1;" class="text">
        <p>${block.value}</p>
    </div>
`;
}
function names(block) {
    return `
    <div style="flex: 0.9;" class="text">
        <p>${block.value}</p>
    </div>
`;
}

function medium(block) {
    return `
    <div style="flex: 1.2;" class="text">
        <p>${block.value}</p>
    </div>
`;
}
function surname(block) {
    return `
    <div style="flex: 1.1;" class="text">
        <p>${block.value}</p>
    </div>
`;
}

function long(block) {
    return `
    <div style="flex: 1.3;  margin-right:15px;" class="text">
        <p>${block.value}</p>
    </div>
`;
}

function action(block) {
    return `
    <div style="flex: 0.6;" class="action">
        <p>${block.value}</p>
    </div>
`;
}
