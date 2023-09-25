import api from "./api";

const URLS = {
  attrsUrl: "/references/attribute",
  operateAttrsUrl: "/references/attribute/dml",
  attrsValues: "/references/attribute/values",
  setDMLValues: "/references/attribute/values/dml",
  createAttributeValues: "/references/attribute/values/create",
};

export const getAttrs = (payload) => {
  return api.post(URLS.attrsUrl, payload);
};

export const operateAttrs = (payload) => {
  return api.post(URLS.operateAttrsUrl, payload);
};

export const getAttrValues = (body) => {
  return api.post(URLS.attrsValues, body);
};

export const setDMLValues = (body) => {
  return api.post(URLS.setDMLValues, body);
};

export const createAttributeValues = (body) => {
  return api.post(URLS.createAttributeValues, body);
};


