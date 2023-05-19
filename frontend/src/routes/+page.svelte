<script lang="ts">
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';
	const stocksData = writable({});
	const selectedStock = writable('');
	const options = writable({});

	$: stocksData.subscribe((data) => {
		options.set({
			chart: {
				type: 'candlestick'
			},
			series: [
				{
					name: $selectedStock,
					data: data[$selectedStock]?.values.reverse().map((v) => {
						return {
							x: new Date(v.datetime),
							y: [v.open, v.high, v.low, v.close]
						};
					})
				}
			]
		});
	});

	function updateOptions(stock: string) {
		const chart = new ApexCharts(document.querySelector('#chart'), $options);
		chart.render();
		selectedStock.set(stock);
		options.update((o) => {
			return {
				...o,
				series: [
					{
						name: stock,
						data: $stocksData[stock]?.values.reverse().map((v) => {
							return {
								x: new Date(v.datetime),
								y: [v.open, v.high, v.low, v.close]
							};
						})
					}
				]
			};
		});
		chart.updateOptions($options);
	}

	onMount(async () => {
		await fetch('http://localhost:8080/api/v1/users/stocks/daily/1')
			.then((res) => res.json())
			.then((data) => {
				selectedStock.set(Object.keys(data)[0]);
				stocksData.set(data);
			});

		const chart = new ApexCharts(document.querySelector('#chart'), $options);
		chart.render();
	});
</script>

<div class="flex flex-row justify-between w-screen h-screen">
	<div class="flex flex-col w-2/12 h-full">
		{#each Object.keys($stocksData) as stock}
			<button on:click={() => updateOptions(stock)}>{stock}</button>
		{/each}
	</div>
	<div class="flex flex-col w-8/12 h-full">
		<div id="chart" class="w-full m-9" />
	</div>
	<div class="flex flex-col w-3/12">news and stuff</div>
</div>
