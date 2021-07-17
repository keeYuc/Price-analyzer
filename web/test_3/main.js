import init, { run_app } from './pkg/test_3';
async function main() {
   await init('/pkg/test_3.wasm');
   run_app();
}
main()