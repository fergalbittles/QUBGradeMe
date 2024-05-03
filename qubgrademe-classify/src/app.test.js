const request = require("supertest");
const app = require('./app.js');

describe("GET /", () => {
  describe("should be a successful request", () => {
    test("should be a 200 status code", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_2: "70", mark_3: "87", mark_4: "65", mark_5: "96" });
  
      expect(response.statusCode).toBe(200);
    });

    test("should be json content type", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_2: "70", mark_3: "87", mark_4: "65", mark_5: "96" });
  
      expect(response.headers["content-type"]).toEqual(expect.stringContaining("json"));
    });

    test("should return error value as false", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_2: "70", mark_3: "87", mark_4: "65", mark_5: "96" });
  
      expect(response.body.error).toEqual(false);
    });

    test("should return correct string value", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_2: "70", mark_3: "87", mark_4: "65", mark_5: "96" });
  
      expect(response.body.string).toEqual("Your overall classification is Upper Second-Class Honours (2:1)");
    });
  });

  describe("should fail due to missing query param", () => {
    test("should be a 400 status code", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_3: "87", mark_4: "65", mark_5: "96" });
  
      expect(response.statusCode).toBe(400);
    });

    test("should be json content type", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_3: "87", mark_4: "65", mark_5: "96" });
  
      expect(response.headers["content-type"]).toEqual(expect.stringContaining("json"));
    });

    test("should return error value as true", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_3: "87", mark_4: "65", mark_5: "96" });
  
      expect(response.body.error).toEqual(true);
    });

    test("should return correct string value", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_3: "87", mark_4: "65", mark_5: "96" });
  
      expect(response.body.string).toEqual("Mark 2 value is missing");
    });
  });

  describe("should fail due to no query params", () => {
    test("should be a 400 status code", async () => {
      const response = await request(app).get("/");
  
      expect(response.statusCode).toBe(400);
    });

    test("should be json content type", async () => {
      const response = await request(app).get("/");
  
      expect(response.headers["content-type"]).toEqual(expect.stringContaining("json"));
    });

    test("should return error value as true", async () => {
      const response = await request(app).get("/");
  
      expect(response.body.error).toEqual(true);
    });

    test("should return correct string value", async () => {
      const response = await request(app).get("/");
  
      expect(response.body.string).toContain("value is missing");
    });
  });

  describe("should fail due to invalid input", () => {
    test("should be a 400 status code", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_2: "87", mark_3: "87", mark_4: "asdasdasd", mark_5: "96" });
  
      expect(response.statusCode).toBe(400);
    });

    test("should be json content type", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_2: "87", mark_3: "87", mark_4: "asdasdasd", mark_5: "96" });
  
      expect(response.headers["content-type"]).toEqual(expect.stringContaining("json"));
    });

    test("should return error value as true", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_2: "87", mark_3: "87", mark_4: "asdasdasd", mark_5: "96" });
  
      expect(response.body.error).toEqual(true);
    });

    test("should return correct string value", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_2: "87", mark_3: "87", mark_4: "asdasdasd", mark_5: "96" });
  
      expect(response.body.string).toEqual("You must provide a valid integer for Mark 4");
    });
  });

  describe("should fail due to negative number input", () => {
    test("should be a 400 status code", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_2: "70", mark_3: "-14", mark_4: "65", mark_5: "96" });
  
      expect(response.statusCode).toBe(400);
    });

    test("should be json content type", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_2: "70", mark_3: "-14", mark_4: "65", mark_5: "96" });
  
      expect(response.headers["content-type"]).toEqual(expect.stringContaining("json"));
    });

    test("should return error value as true", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_2: "70", mark_3: "-14", mark_4: "65", mark_5: "96" });
  
      expect(response.body.error).toEqual(true);
    });

    test("should return correct string value", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_2: "70", mark_3: "-14", mark_4: "65", mark_5: "96" });
  
      expect(response.body.string).toEqual("You must provide a non-negative integer for Mark 3");
    });
  });

  describe("should fail due to value over 100", () => {
    test("should be a 400 status code", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_2: "7912", mark_3: "14", mark_4: "65", mark_5: "96" });
  
      expect(response.statusCode).toBe(400);
    });

    test("should be json content type", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_2: "7912", mark_3: "14", mark_4: "65", mark_5: "96" });
  
      expect(response.headers["content-type"]).toEqual(expect.stringContaining("json"));
    });

    test("should return error value as true", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_2: "7912", mark_3: "14", mark_4: "65", mark_5: "96" });
  
      expect(response.body.error).toEqual(true);
    });

    test("should return correct string value", async () => {
      const response = await request(app).get("/").query({ mark_1: "30", mark_2: "7912", mark_3: "14", mark_4: "65", mark_5: "96" });
  
      expect(response.body.string).toEqual("You cannot exceed 100 marks for Mark 2");
    });
  });
});