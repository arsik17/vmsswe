let responseBody = "";

async function getAllUsers() {
    const response = await fetch("http://localhost:8080/users/all", {
        headers: {
            "content-type": "application/json",
            accept: "application/json",
        },
    });

    responseBody = await response.json();

    console.log(responseBody);

    const $list = document.querySelector(".list");
    let elements = "";
    responseBody.forEach((oneinstance) => {
        elements += `
        <div class="element">
            <div style="width: 50px;" class="text"  id="user_id">
                <p>${oneinstance.user_id}</p>
            </div>
            <div style="flex: 1.7;" class="text">
                <p>${oneinstance.email}</p>
            </div>
            <div style="flex: 1.2;" class="text">
                <p>${oneinstance.given_name}</p>
            </div>
            <div style="flex: 1.2;" class="text">
                <p>${oneinstance.surname}</p>
            </div>
            <div style="flex: 1.2;" class="text">
                <p>${oneinstance.phone_number}</p>
            </div>
            <div style="flex: 0.8;" class="text">
                <p>${oneinstance.city}</p>
            </div>
            <div class = "change-bt">
            
            <button onclick="deleteFunction(this)" type="button" class="change" id="d">
                <img
                    src="/assets/bin.png"
                    alt=""
                />
                </button>
            <button onclick="viewFunction(this)" type="button" class="change" id="c">
                <img
                    src="/assets/refresh.png"
                    alt=""
                />
                    </button>
            </div>
            
        </div>`;
        $list.innerHTML = elements;
    });
}
getAllUsers();

function element(id) {
    return `<div class="element" id="${id}">
    </div>
    `;
}

function id(content) {
    return `
        <div class="text">
            <p>${content}</p>
        </div>
    `;
}
function email(content) {
    return `
        <div class="text">
            <p>${content}</p>
        </div>
    `;
}
function given_name(content) {
    return `
        <div class="text">
            <p>${content}</p>
        </div>
    `;
}
function surname(content) {
    return `
        <div class="text">
            <p>${content}</p>
        </div>
    `;
}

function city(content) {
    return `
        <div class="text">
            <p>${content}</p>
        </div>
    `;
}

function phone_number(content) {
    return `
        <div class="text">
            <p>${content}</p>
        </div>
    `;
}

function change() {
    return `
        <button type="button" id="change">
        <img
            src="/assets/refresh.png"
            alt=""
        />
            </button>
    `;
}
