import Ranking from './Ranking';
import { render, screen } from '@testing-library/react';
import { act } from "react-dom/test-utils";
import axios from 'axios'
import { Routes, Route, BrowserRouter, Router } from "react-router-dom";

test('test Ranking component', async () => {
  jest.spyOn(axios, 'get').mockResolvedValue({"data": {
    "players": [
      {
        "name": "福谷　浩司",
        "rate": 1560,
        "id": 2007
      },
      {
        "name": "鈴木　誠也",
        "rate": 1545,
        "id": 2010
      },
      {
        "name": "福地　和広",
        "rate": 1530,
        "id": 1539
      },
      {
        "name": "岩隈　久志",
        "rate": 1515,
        "id": 1635
      }
    ]
  }});

  render(
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Ranking />} />
      </Routes>
    </BrowserRouter>
  )

  //[TODO] I want to increase test coverage...
  expect(await screen.findByText(/福谷.+浩司/)).toBeInTheDocument();
  expect(await screen.findByText(/1545/)).toBeInTheDocument();
  expect(await screen.findByText(/岩隈.+久志/)).toBeInTheDocument();
  expect(await screen.findByText(/1515/)).toBeInTheDocument();
});