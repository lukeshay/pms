import { FC, ReactNode } from "react";

export const BaseLayout: FC<{ children: ReactNode }> = ({ children }) => (
	<div className="flex min-h-screen justify-center bg-base-200 p-6">
		<div className="w-full max-w-6xl">
			<nav className="navbar rounded-box bg-base-100">
				<div className="flex-1">
					<a className="btn-ghost btn text-xl normal-case">daisyUI</a>
				</div>
				<div className="flex-none">
					<ul className="menu menu-horizontal px-1">
						<li>
							<a>Link</a>
						</li>
						<li>
							<details>
								<summary>Parent</summary>
								<ul className="bg-base-100 p-2">
									<li>
										<a>Link 1</a>
									</li>
									<li>
										<a>Link 2</a>
									</li>
								</ul>
							</details>
						</li>
					</ul>
				</div>
			</nav>
			<section id="content" className="pb-16 pt-8">
				{children}
			</section>
		</div>
	</div>
);
