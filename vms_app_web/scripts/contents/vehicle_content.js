async function getAllUsers() {
    const response = await fetch("/vehicles", {
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
            <div style="flex: 1.3;" class="text"  id="user_id">
                <p>${oneinstance.vin}</p>
            </div>
            <div style="flex: 0.7;" class="text">
                <p>${oneinstance.license_plate}</p>
            </div>
            <div style="flex: 0.9;" class="text">
                <p>${oneinstance.car_make}</p>
            </div>
            <div style="flex: 1;" class="text">
                <p>${oneinstance.model}</p>
            </div>
            <div style="flex: 0.5;" class="text">
                <p>${oneinstance.production_year}</p>
            </div>
            <div style="flex: 0.65;" class="text">
                <p>${oneinstance.mileage}</p>
            </div>
            <div style="flex: 0.7;" class="text">
                <p>${oneinstance.activity_status}</p>
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
