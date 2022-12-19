import { defineCustomElement, h } from "vue";
import SearchBar from "./SearchBar.vue";

console.log(SearchBar.styles);

const SearchBarElement = defineCustomElement(SearchBar);

customElements.define("search-bar", SearchBarElement);
