import React, { useState, useEffect } from 'react';
import axios from 'axios';

const TugasAxios = () => {
    const [books, setBooks] = useState([]);
    const [form, setForm] = useState({
        title: '',
        description: '',
        image_url: '',
        release_year: '',
        price: '',
        total_page: '',
    });
    const [errors, setErrors] = useState({});
    const [editing, setEditing] = useState(null);

    useEffect(() => {
        fetchBooks();
    }, []);

    const fetchBooks = async () => {
        try {
            const response = await axios.get('http://localhost:8080/books');
            setBooks(response.data);
        } catch (error) {
            console.error('Error fetching books:', error);
        }
    };

    const handleInputChange = (e) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value,
        });
    };

    const validateForm = () => {
        const { title, description, image_url, release_year, price, total_page } = form;
        let errors = {};

        if (!title) errors.title = 'Title is required';
        if (!description) errors.description = 'Description is required';
        if (!image_url) {
            errors.image_url = 'Image URL is required';
        } else if (!/^https?:\/\/.+\..+/.test(image_url)) {
            errors.image_url = 'Image URL must be a valid URL';
        }
        if (!release_year) {
            errors.release_year = 'Release year is required';
        } else if (release_year < 1980 || release_year > 2021) {
            errors.release_year = 'Release year must be between 1980 and 2021';
        }
        if (!price) errors.price = 'Price is required';
        if (!total_page) {
            errors.total_page = 'Total page is required';
        } else if (isNaN(total_page) || total_page <= 0) {
            errors.total_page = 'Total page must be a positive number';
        }

        setErrors(errors);
        return Object.keys(errors).length === 0;
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        if (!validateForm()) return;

        try {
            const config = {
                headers: {
                    'Content-Type': 'application/json',
                },
            };

            console.log("Submitting form data: ", form);

            if (editing) {
                await axios.patch(`http://localhost:8080/books/${editing}`, JSON.stringify(form), config);
            } else {
                await axios.post('http://localhost:8080/books', JSON.stringify(form), config);
            }


            setForm({
                title: '',
                description: '',
                image_url: '',
                release_year: '',
                price: '',
                total_page: '',
                thickness: '',
            });
            setEditing(null);
            fetchBooks();
        } catch (error) {
            console.error('Error submitting form:', error);
        }
    };

    const handleEdit = (book) => {
        setForm({
            title: book.title,
            description: book.description,
            image_url: book.image_url,
            release_year: book.release_year,
            price: book.price,
            total_page: book.total_page,
        });
        setEditing(book.id);
    };

    const handleDelete = async (id) => {
        try {
            await axios.delete(`http://localhost:8080/books/${id}`);
            fetchBooks();
        } catch (error) {
            console.error('Error deleting book:', error);
        }
    };

    return (
        <div>
            <h2>Book List</h2>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Title:</label>
                    <input type="text" name="title" value={form.title} onChange={handleInputChange} />
                    {errors.title && <p>{errors.title}</p>}
                </div>
                <div>
                    <label>Description:</label>
                    <input type="text" name="description" value={form.description} onChange={handleInputChange} />
                    {errors.description && <p>{errors.description}</p>}
                </div>
                <div>
                    <label>Image URL:</label>
                    <input type="text" name="image_url" value={form.image_url} onChange={handleInputChange} />
                    {errors.image_url && <p>{errors.image_url}</p>}
                </div>
                <div>
                    <label>Release Year:</label>
                    <input type="number" name="release_year" value={form.release_year} onChange={handleInputChange} />
                    {errors.release_year && <p>{errors.release_year}</p>}
                </div>
                <div>
                    <label>Price:</label>
                    <input type="text" name="price" value={form.price} onChange={handleInputChange} />
                    {errors.price && <p>{errors.price}</p>}
                </div>
                <div>
                    <label>Total Page:</label>
                    <input type="number" name="total_page" value={form.total_page} onChange={handleInputChange} />
                    {errors.total_page && <p>{errors.total_page}</p>}
                </div>
                <button type="submit">Submit</button>
            </form>

            <table>
                <thead>
                <tr>
                    <th>Title</th>
                    <th>Description</th>
                    <th>Image</th>
                    <th>Release Year</th>
                    <th>Price</th>
                    <th>Total Page</th>
                    <th>Actions</th>
                </tr>
                </thead>
                <tbody>
                {books.map(book => (
                    <tr key={book.id}>
                        <td>{book.title}</td>
                        <td>{book.description}</td>
                        <td><img src={book.image_url} alt={book.title} width="50" /></td>
                        <td>{book.release_year}</td>
                        <td>{book.price}</td>
                        <td>{book.total_page}</td>
                        <td>
                            <button onClick={() => handleEdit(book)}>Edit</button>
                            <button onClick={() => handleDelete(book.id)}>Delete</button>
                        </td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};

export default TugasAxios;
