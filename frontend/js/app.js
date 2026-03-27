const API = "/api";

window.onload = () => {
    loadSkills();
    loadDashboard();
};

// Load Skills
async function loadSkills() {
    const res = await fetch(`${API}/skills`);
    const skills = await res.json();

    const container = document.getElementById("skills");
    container.innerHTML = "";

    skills.forEach(skill => {
        const percent = skill.target_hours > 0
            ? Math.min((skill.total_hours / skill.target_hours) * 100, 100)
            : 0;

        const div = document.createElement("div");
        div.className = "skill";

        div.innerHTML = `
            <div class="skill-header">
                <strong>${skill.name}</strong>
                <span>${skill.total_hours} / ${skill.target_hours} hrs</span>
            </div>
            <small>${skill.category}</small>

            <div class="progress">
                <div class="progress-bar" style="width:${percent}%"></div>
            </div>

            <div class="actions">
                <button class="log" onclick="logHours(${skill.id})">+ Log</button>
                <button class="delete" onclick="deleteSkill(${skill.id})">Delete</button>
            </div>
        `;

        container.appendChild(div);
    });
}

// Add Skill
async function addSkill() {
    const name = document.getElementById("name").value;
    const category = document.getElementById("category").value;
    const target_hours = parseInt(document.getElementById("target_hours").value);

    if (!name) return alert("Skill name required");

    await fetch(`${API}/skills`, {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({ name, category, target_hours })
    });

    clearInputs();
    loadSkills();
    loadDashboard();
}

// Delete Skill
async function deleteSkill(id) {
    if (!confirm("Delete this skill?")) return;

    await fetch(`${API}/skills/${id}`, {
        method: "DELETE"
    });

    loadSkills();
    loadDashboard();
}

// Log Hours
async function logHours(id) {
    const hours = prompt("Enter hours:");
    if (!hours) return;

    await fetch(`${API}/skills/${id}/log`, {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({
            hours: parseFloat(hours),
            notes: "",
            log_date: new Date().toISOString().split("T")[0]
        })
    });

    loadSkills();
    loadDashboard();
}

// Dashboard
async function loadDashboard() {
    const res = await fetch(`${API}/dashboard`);
    const data = await res.json();

    document.getElementById("dashboard").innerHTML = `
        <p><strong>Total Skills:</strong> ${data.total_skills}</p>
        <p><strong>Total Hours:</strong> ${data.total_hours}</p>
    `;
}

// Clear inputs
function clearInputs() {
    document.getElementById("name").value = "";
    document.getElementById("category").value = "";
    document.getElementById("target_hours").value = "";
}
