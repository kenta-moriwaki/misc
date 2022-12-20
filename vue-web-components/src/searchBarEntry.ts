import { defineCustomElement } from "vue";
import SearchBar from "./SearchBar.vue";

const SearchBarElement = defineCustomElement(SearchBar);

customElements.define("search-bar", SearchBarElement);
