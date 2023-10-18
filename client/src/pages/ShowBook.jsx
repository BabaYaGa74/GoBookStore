import React, { useState, useEffect } from "react";
import axios from "axios";
import { useParams, Link } from "react-router-dom";
import BackButton from "../components/BackButton";
import Spinner from "../components/Spinner";

const ShowBook = () => {
  const [book, setBooks] = useState({});
  const [loading, setLoading] = useState();
  const { bookId } = useParams();
  useEffect(() => {
    setLoading(true);
    axios
      .get(`http://localhost:5000/book/${bookId}`)
      .then((response) => {
        setBooks(response.data);
        setLoading(false);
      })
      .catch((error) => {
        console.log(error);
        setLoading(false);
      });
  }, []);
  return (
    <div className="p-4">
      <BackButton />
      <h1 className="text-3xl my-4">Show Book</h1>
      {loading ? (
        <Spinner />
      ) : (
        <div className="flex flex-col border-2 border-sky-400 rounded-xl w-fit p-4">
          <div className="my-4">
            <span className="text-xl mr-4 text-gray-500">Id :</span>
            <span>{book.ID}</span>
          </div>
          <div className="my-4">
            <span className="text-xl mr-4 text-gray-500">Title :</span>
            <span>{book.name}</span>
          </div>
          <div className="my-4">
            <span className="text-xl mr-4 text-gray-500">Author :</span>
            <span>{book.author}</span>
          </div>
          <div className="my-4">
            <span className="text-xl mr-4 text-gray-500"> Publication :</span>
            <span>{book.publication}</span>
          </div>
        </div>
      )}
    </div>
  );
};

export default ShowBook;
