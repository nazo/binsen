import test from 'ava';
import assert from 'power-assert';
import { State, mutations } from './index';

it('test mutations setLang', () => {
  const state: State = new State();
  assert.equal(state.locale, 'en');
  mutations.setLang(state, 'ja');
  assert.equal(state.locale, 'ja');
});
