import { createApp } from 'https://unpkg.com/vue@3/dist/vue.esm-browser.js';

const App = {
  data() {
    return { items: [], loading: true, error: '' };
  },
  async mounted() {
    await this.load();
  },
  methods: {
    async load() {
      this.loading = true; this.error = '';
      try {
        const res = await fetch('/api/items');
        if (!res.ok) throw new Error(`HTTP ${res.status}`);
        this.items = await res.json();
      } catch (e) {
        this.error = 'Failed to load items. Please retry.';
      } finally {
        this.loading = false;
      }
    }
  },
  template: `
    <div>
      <div v-if="error" class="text-red-600 mb-3">
        {{ error }} <button @click="load" class="underline">Retry</button>
      </div>
      <div v-if="loading" class="flex items-center gap-2 text-gray-500 mb-3">
        <svg class="animate-spin h-5 w-5 text-gray-500" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v4a4 4 0 00-4 4H4z"></path></svg>
        <span>Loading items…</span>
      </div>
      <div v-else>
        <div v-if="items.length === 0" class="text-gray-500">No items yet. Add one to get started.</div>
        <ul v-else class="grid gap-3">
          <li v-for="it in items" :key="it.id" class="bg-white rounded shadow p-3 flex gap-3 items-center">
            <img :src="it.thumbnailUrl" alt="thumbnail" class="w-16 h-16 rounded object-cover bg-gray-100" />
            <div class="flex-1">
              <div class="font-medium">{{ it.name }}</div>
              <div>
                <span :class="{
                  'inline-block text-xs px-2 py-0.5 rounded bg-emerald-100 text-emerald-800': it.availability==='available',
                  'inline-block text-xs px-2 py-0.5 rounded bg-gray-200 text-gray-700': it.availability!=='available'
                }">{{ it.availability }}</span>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>
  `
};

createApp(App).mount('#list');