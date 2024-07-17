import React from 'react';

const thingsToDo = [
    "Belajar GIT & CLI",
    "Belajar HTML & CSS",
    "Belajar Javascript",
    "Belajar ReactJS Dasar",
    "Belajar ReactJS Advance"
];

const CheckboxList = ({ items }) => (
    <div className="checkbox-list">
        {items.map((item, index) => (
            <label key={index} className="checkbox-item">
                <input type="checkbox" id={`item-${index}`} className="checkbox-input" />
                {item}
            </label>
        ))}
    </div>
);

function TugasIntroReact() {
    return (
        <div>
            <CheckboxList items={thingsToDo} />
            <button className="send-button">SEND</button>
        </div>
    );
}

export default TugasIntroReact;
