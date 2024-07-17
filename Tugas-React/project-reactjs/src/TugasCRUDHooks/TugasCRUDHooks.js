import React, { useState } from 'react';

const initialData = [
    { nama: "Nanas", hargaTotal: 100000, beratTotal: 4000 },
    { nama: "Manggis", hargaTotal: 350000, beratTotal: 10000 },
    { nama: "Nangka", hargaTotal: 90000, beratTotal: 2000 },
    { nama: "Durian", hargaTotal: 400000, beratTotal: 5000 },
    { nama: "Strawberry", hargaTotal: 120000, beratTotal: 6000 }
];

const TugasCRUDHooks = () => {
    const [data, setData] = useState(initialData);
    const [form, setForm] = useState({ nama: '', hargaTotal: '', beratTotal: '' });
    const [editingIndex, setEditingIndex] = useState(null);

    const handleInputChange = (e) => {
        const { name, value } = e.target;
        setForm({ ...form, [name]: value });
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        if (form.nama && form.hargaTotal && form.beratTotal) {
            if (editingIndex !== null) {
                const updatedData = [...data];
                updatedData[editingIndex] = form;
                setData(updatedData);
                setEditingIndex(null);
            } else {
                setData([...data, form]);
            }
            setForm({ nama: '', hargaTotal: '', beratTotal: '' });
        }
    };

    const handleEdit = (index) => {
        setForm(data[index]);
        setEditingIndex(index);
    };

    const handleDelete = (index) => {
        const updatedData = data.filter((_, i) => i !== index);
        setData(updatedData);
    };

    return (
        <div>
            <h1>Daftar Harga Buah</h1>
            <table border="1">
                <thead>
                <tr>
                    <th>No</th>
                    <th>Nama</th>
                    <th>Harga total</th>
                    <th>Berat total</th>
                    <th>Harga per kg</th>
                    <th>Aksi</th>
                </tr>
                </thead>
                <tbody>
                {data.map((item, index) => (
                    <tr key={index}>
                        <td>{index + 1}</td>
                        <td>{item.nama}</td>
                        <td>{item.hargaTotal}</td>
                        <td>{item.beratTotal / 1000} kg</td>
                        <td>{item.hargaTotal / (item.beratTotal / 1000)}</td>
                        <td>
                            <button onClick={() => handleEdit(index)}>Edit</button>
                            <button onClick={() => handleDelete(index)}>Delete</button>
                        </td>
                    </tr>
                ))}
                </tbody>
            </table>
            <h1>Form Daftar Harga Buah</h1>
            <form onSubmit={handleSubmit}>
                <label>
                    Nama:
                    <input type="text" name="nama" value={form.nama} onChange={handleInputChange} />
                </label>
                <br />
                <label>
                    Harga Total:
                    <input type="number" name="hargaTotal" value={form.hargaTotal} onChange={handleInputChange} />
                </label>
                <br />
                <label>
                    Berat Total (dalam gram):
                    <input type="number" name="beratTotal" value={form.beratTotal} onChange={handleInputChange} />
                </label>
                <br />
                <button type="submit">Submit</button>
            </form>
        </div>
    );
};

export default TugasCRUDHooks;
