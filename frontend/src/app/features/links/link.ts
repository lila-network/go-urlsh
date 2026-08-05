export interface LinksResponse {
  data: Link[];
  status: string;
}

export interface LinkResponse {
  data: Link;
  status: string;
}

export interface Link {
  id: string
  url: string
  description: string
  created: string
  modified: string
}

export interface LinkCreationRequest {
  id: string;
  url: string;
  description: string;
}

export interface LinksErrorResponse {
  status: string
  message: string
}
