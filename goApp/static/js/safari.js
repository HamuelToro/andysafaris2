import { safaris } from "./data/safari.js";

const daysBtns = document.querySelectorAll(".days-options .day");
daysBtns.forEach((btn) => {
	btn.addEventListener("click", () => {
		const day = Number(btn.dataset.dayValue);

		if (document.querySelector(".day-selected")) {
			document.querySelector(".day-selected").classList.remove("day-selected");
		}

		btn.classList.add("day-selected");

		filterSafari(day, btn);
	});
});

document.querySelector(".show-all").addEventListener("click", () => {
	if (document.querySelector(".day-selected")) {
		document.querySelector(".day-selected").classList.remove("day-selected");
	}

	filterSafari(100);
});

function filterSafari(day) {
	const safariContainer = document.getElementById("js-safari-grid");
	safariContainer.innerHTML = "";

	const filteredSafaris =
		day === 100 ? safaris : safaris.filter((safari) => safari.Days === day);

	filteredSafaris.forEach((safari) => {
		safariContainer.innerHTML += `
            <div class="safari-option">
                <img class="safari-image" src="${safari.Image}" alt="" />
                <p class="section-title safari-card-title">
                    ${safari.Days} Days ${safari.Title}
                    <span class="last-two-words">By ${safari.Guide}</span>
                </p>
                <p class="safari-card-description">
                    ${safari.Description}
                </p>
                <div class="safari-learn-more btn" onclick="window.location.href='/safaris/${safari.ID}'">Learn More</div>
            </div>
        `;
	});
}
