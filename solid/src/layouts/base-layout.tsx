import { A } from "@solidjs/router";
import { JSX, ParentComponent, Show } from "solid-js";

import { useGlobalStore } from "../provider";

export type BaseLayoutProperties = {
	onSplitPanelClose?: () => void;
	splitPanel?: JSX.Element;
	title?: string;
};

export const BaseLayout: ParentComponent<BaseLayoutProperties> = (properties) => {
	const { setClaims, setTokens, state } = useGlobalStore();

	return (
		<div class="flex min-h-screen bg-base-200">
			<div class="flex w-full justify-center px-3 py-6">
				<div class="w-full max-w-6xl">
					<nav class="navbar rounded-box bg-base-100">
						<div class="flex-1">
							<A class="btn-ghost btn text-xl normal-case" href="/">
								{"PMS"}
							</A>
						</div>
						<div class="flex-none">
							<ul class="menu menu-horizontal px-1">
								{state.claims && (
									<li>
										<A
											classList={{
												active: window.location.pathname === "/",
											}}
											href="/"
										>
											{"Books"}
										</A>
									</li>
								)}
								<li>
									<details>
										<summary>Profile</summary>
										<ul class="bg-base-100 p-2">
											{state.claims ? (
												<>
													<li>
														<A href="/profile">{"Profile"}</A>
													</li>
													<li>
														<button
															onClick={() => {
																setTokens();
																setClaims();
																// navigate("/sign-in");
															}}
														>
															{"Sign Out"}
														</button>
													</li>
												</>
											) : (
												<li>
													<A href="/sign-in">{"Sign In"}</A>
												</li>
											)}
										</ul>
									</details>
								</li>
							</ul>
						</div>
					</nav>
					<section class="pb-16 pt-8" id="content">
						{properties.children}
					</section>
				</div>
			</div>
			<Show when={properties.splitPanel}>{properties.splitPanel}</Show>
		</div>
	);
};
