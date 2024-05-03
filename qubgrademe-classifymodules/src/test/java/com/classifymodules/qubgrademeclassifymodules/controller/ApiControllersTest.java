package com.classifymodules.qubgrademeclassifymodules.controller;

import com.classifymodules.qubgrademeclassifymodules.util.ClassifyModules;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.junit.jupiter.SpringExtension;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.MvcResult;
import org.springframework.test.web.servlet.RequestBuilder;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;

import java.util.HashMap;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.when;

@ExtendWith(SpringExtension.class)
@WebMvcTest(ApiControllers.class)
class ApiControllersTest {

    @Autowired
    private MockMvc mvc;

    @MockBean
    private ClassifyModules classifyModules;

    @Test
    void getPage_successfulRequest() throws Exception {
        HashMap<String, Object> parsedMarks = new HashMap<>();
        HashMap<String, Object> classifiedMarks = new HashMap<>();

        when(classifyModules.classifyModules(parsedMarks)).thenReturn(classifiedMarks);

        RequestBuilder request = MockMvcRequestBuilders
                .get("/")
                .param("module_1", "Databases")
                .param("module_2", "Concurrent Programming")
                .param("module_3", "Cloud Computing")
                .param("module_4", "Cyber Security Fundamentals")
                .param("module_5", "Object Oriented Programming")
                .param("mark_1", "92")
                .param("mark_2", "77")
                .param("mark_3", "81")
                .param("mark_4", "75")
                .param("mark_5", "68");

        MvcResult response = mvc.perform(request).andReturn();

        assertEquals(200, response.getResponse().getStatus());
        assertTrue(response.getResponse().getContentType().equals("application/json"));
        assertTrue(response.getResponse().getContentAsString().contains("\"error\":false"));
    }

    @Test
    void getPage_missingParam() throws Exception {
        RequestBuilder request = MockMvcRequestBuilders
                .get("/")
                .param("module_1", "Databases")
                .param("module_2", "Concurrent Programming")
                .param("module_4", "Cyber Security Fundamentals")
                .param("module_5", "Object Oriented Programming")
                .param("mark_1", "92")
                .param("mark_2", "77")
                .param("mark_3", "81")
                .param("mark_4", "75")
                .param("mark_5", "68");

        MvcResult response = mvc.perform(request).andReturn();

        assertEquals(400, response.getResponse().getStatus());
        assertTrue(response.getResponse().getContentType().equals("application/json"));
        assertTrue(response.getResponse().getContentAsString().contains("\"error\":true"));
        assertTrue(response.getResponse().getContentAsString().contains("\"string\":\"Module 3 value is missing\""));
    }

    @Test
    void getPage_noParams() throws Exception {
        RequestBuilder request = MockMvcRequestBuilders.get("/");

        MvcResult response = mvc.perform(request).andReturn();

        assertEquals(400, response.getResponse().getStatus());
        assertTrue(response.getResponse().getContentType().equals("application/json"));
        assertTrue(response.getResponse().getContentAsString().contains("\"error\":true"));
        assertTrue(response.getResponse().getContentAsString().contains("value is missing\""));
    }

    @Test
    void getPage_invalidInput() throws Exception {
        RequestBuilder request = MockMvcRequestBuilders
                .get("/")
                .param("module_1", "Databases")
                .param("module_2", "Concurrent Programming")
                .param("module_3", "Cloud Computing")
                .param("module_4", "Cyber Security Fundamentals")
                .param("module_5", "Object Oriented Programming")
                .param("mark_1", "sdfsdfsdf")
                .param("mark_2", "77")
                .param("mark_3", "81")
                .param("mark_4", "75")
                .param("mark_5", "68");

        MvcResult response = mvc.perform(request).andReturn();

        assertEquals(400, response.getResponse().getStatus());
        assertTrue(response.getResponse().getContentType().equals("application/json"));
        assertTrue(response.getResponse().getContentAsString().contains("\"error\":true"));
        assertTrue(response.getResponse().getContentAsString().contains("\"string\":\"You must provide a valid integer for Mark 1\""));
    }

    @Test
    void getPage_negativeValue() throws Exception {
        RequestBuilder request = MockMvcRequestBuilders
                .get("/")
                .param("module_1", "Databases")
                .param("module_2", "Concurrent Programming")
                .param("module_3", "Cloud Computing")
                .param("module_4", "Cyber Security Fundamentals")
                .param("module_5", "Object Oriented Programming")
                .param("mark_1", "92")
                .param("mark_2", "77")
                .param("mark_3", "81")
                .param("mark_4", "-14")
                .param("mark_5", "68");

        MvcResult response = mvc.perform(request).andReturn();

        assertEquals(400, response.getResponse().getStatus());
        assertTrue(response.getResponse().getContentType().equals("application/json"));
        assertTrue(response.getResponse().getContentAsString().contains("\"error\":true"));
        assertTrue(response.getResponse().getContentAsString().contains("\"string\":\"You must provide a non-negative integer for Mark 4\""));
    }

    @Test
    void getPage_exceededMaxValue() throws Exception {
        RequestBuilder request = MockMvcRequestBuilders
                .get("/")
                .param("module_1", "Databases")
                .param("module_2", "Concurrent Programming")
                .param("module_3", "Cloud Computing")
                .param("module_4", "Cyber Security Fundamentals")
                .param("module_5", "Object Oriented Programming")
                .param("mark_1", "92")
                .param("mark_2", "9898")
                .param("mark_3", "81")
                .param("mark_4", "75")
                .param("mark_5", "68");

        MvcResult response = mvc.perform(request).andReturn();

        assertEquals(400, response.getResponse().getStatus());
        assertTrue(response.getResponse().getContentType().equals("application/json"));
        assertTrue(response.getResponse().getContentAsString().contains("\"error\":true"));
        assertTrue(response.getResponse().getContentAsString().contains("\"string\":\"You cannot exceed 100 marks for Mark 2\""));
    }

    @Test
    void getPage_emptyModuleValue() throws Exception {
        RequestBuilder request = MockMvcRequestBuilders
                .get("/")
                .param("module_1", "Databases")
                .param("module_2", "")
                .param("module_3", "Cloud Computing")
                .param("module_4", "Cyber Security Fundamentals")
                .param("module_5", "Object Oriented Programming")
                .param("mark_1", "92")
                .param("mark_2", "84")
                .param("mark_3", "81")
                .param("mark_4", "75")
                .param("mark_5", "68");

        MvcResult response = mvc.perform(request).andReturn();

        assertEquals(400, response.getResponse().getStatus());
        assertTrue(response.getResponse().getContentType().equals("application/json"));
        assertTrue(response.getResponse().getContentAsString().contains("\"error\":true"));
        assertTrue(response.getResponse().getContentAsString().contains("\"string\":\"Module 2 value is missing\""));
    }

    @Test
    void getPage_whitespaceModuleValue() throws Exception {
        RequestBuilder request = MockMvcRequestBuilders
                .get("/")
                .param("module_1", "Databases")
                .param("module_2", "      ")
                .param("module_3", "Cloud Computing")
                .param("module_4", "Cyber Security Fundamentals")
                .param("module_5", "Object Oriented Programming")
                .param("mark_1", "92")
                .param("mark_2", "84")
                .param("mark_3", "81")
                .param("mark_4", "75")
                .param("mark_5", "68");

        MvcResult response = mvc.perform(request).andReturn();

        assertEquals(400, response.getResponse().getStatus());
        assertTrue(response.getResponse().getContentType().equals("application/json"));
        assertTrue(response.getResponse().getContentAsString().contains("\"error\":true"));
        assertTrue(response.getResponse().getContentAsString().contains("\"string\":\"Module 2 value is missing\""));
    }

}