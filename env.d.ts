/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly API_URL: string;
  readonly SHAD_UI_COMP: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
