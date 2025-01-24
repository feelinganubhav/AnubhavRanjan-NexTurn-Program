
import { delayedGreeting } from '../src/a3_Async_Functions_Delayed_Greeting';

describe('Async Functions: Delayed Greeting', () => {
  beforeEach(() => {
    jasmine.clock().install();
  });

  afterEach(() => {
    jasmine.clock().uninstall();
  });

  it('should resolve with the correct greeting message', async () => {
    const name = 'Anubhav';
    const delay = 1000;

    const promise = delayedGreeting(name, delay);

    jasmine.clock().tick(delay);

    const result = await promise;

    expect(result).toBe(`Hello, ${name}!`);
  });

  it('should respect the specified delay', async () => {
    const name = 'Anubhav';
    const delay = 2000;

    const promise = delayedGreeting(name, delay);

    let isResolved = false;
    promise.then(() => (isResolved = true));

    jasmine.clock().tick(delay - 500);
    expect(isResolved).toBe(false);

    jasmine.clock().tick(500);

    const result = await promise;
    expect(isResolved).toBe(true);

    expect(result).toBe(`Hello, ${name}!`);
  });

  it('should handle multiple concurrent delays correctly', async () => {
    const promises = [
      delayedGreeting('Anubhav', 1000),
      delayedGreeting('Jonty', 2000),
    ];

    let isResolved1 = false;
    promises[0].then(() => (isResolved1 = true));

    let isResolved2 = false;
    promises[1].then(() => (isResolved2 = true));


    jasmine.clock().tick(1000);

    const firstResult = await promises[0];
    expect(isResolved1).toBe(true);
    expect(firstResult).toBe('Hello, Anubhav!');

    jasmine.clock().tick(1000);
    
    const secondResult = await promises[1];
    expect(isResolved2).toBe(true);
    expect(secondResult).toBe('Hello, Jonty!');
  });

  it('should handle zero delay', async () => {
    const name = 'Anubhav';
    const delay = 0;

    const promise = delayedGreeting(name, delay);

    
    jasmine.clock().tick(0);

    const result = await promise;
    expect(result).toBe(`Hello, ${name}!`);
  });
});
