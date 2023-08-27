import react from '@vitejs/plugin-react-swc';
import * as path from 'path';
import { defineConfig } from 'vite';
import eslint from 'vite-plugin-eslint';
import svgr from 'vite-plugin-svgr';

export default defineConfig({
  plugins: [
    react(),
    svgr(), // To easily load SVGs as React components
    eslint({
      include: ['./src/**/*.ts', './src/**/*.tsx'],
      cache: false,
      lintOnStart: true,
      failOnError: false,
      failOnWarning: false,
    }),
  ],
  envDir: './env',
  resolve: {
    // Replace @ with src
    alias: [{ find: '@', replacement: path.resolve(__dirname, 'src') }],
  },
  server: {
    host: true,
  },
});
